**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-08 17:26:13 UTC（2026-03-09 01:26:13 UTC+8）

**代理总数：55**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 212ms | 否 | ✓ 709ms | ✓ 1049ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1628ms | ✓ 1558ms | ✓ 1347ms | ✓ 1114ms | ✓ 892ms | http |
| 178.236.245.59:3128 | ✓ 1126ms | 否 | ✓ 1700ms | 否 | ✓ 1974ms | http |
| 178.236.245.17:3128 | ✓ 1156ms | 否 | ✓ 1699ms | 否 | ✓ 1949ms | http |
| 217.76.245.80:999 | ✓ 711ms | ✓ 1474ms | ✓ 1083ms | ✓ 1379ms | ✓ 1100ms | http |
| 192.227.137.72:5050 | ✓ 897ms | 否 | ✓ 872ms | ✓ 1149ms | ✓ 1140ms | http |
| 205.209.118.30:3138 | ✓ 870ms | ✓ 1820ms | 否 | ✓ 1074ms | ✓ 793ms | http |
| 152.42.213.210:8080 | ✓ 1560ms | 否 | ✓ 1484ms | ✓ 1284ms | ✓ 1024ms | http |
| 165.227.5.10:8888 | ✓ 322ms | ✓ 1951ms | 否 | ✓ 1091ms | ✓ 1190ms | http |
| 194.213.18.200:443 | ✓ 1370ms | ✓ 1825ms | 否 | ✓ 1797ms | 否 | http |
| 202.155.12.161:443 | ✓ 1510ms | 否 | ✓ 1581ms | ✓ 1261ms | ✓ 1145ms | http |
| 101.43.255.96:80 | ✓ 1190ms | ✓ 1782ms | 否 | 否 | ✓ 1921ms | http |
| 185.115.74.185:8080 | ✓ 1416ms | ✓ 1580ms | ✓ 1345ms | 否 | 否 | http |
| 88.80.150.82:8080 | ✓ 1207ms | 否 | 否 | ✓ 1967ms | ✓ 1802ms | https |
| 116.80.82.228:3172 | ✓ 1724ms | 否 | ✓ 1934ms | 否 | ✓ 1870ms | http |
| 120.92.212.16:8890 | ✓ 1154ms | ✓ 1417ms | ✓ 1528ms | ✓ 1429ms | 否 | http |
| 201.150.116.32:999 | ✓ 1079ms | 否 | ✓ 1259ms | ✓ 1280ms | 否 | http |
| 5.101.0.233:3128 | ✓ 1303ms | 否 | ✓ 1080ms | 否 | ✓ 1728ms | http |
| 152.42.213.210:80 | ✓ 1672ms | 否 | ✓ 1386ms | 否 | ✓ 1220ms | http |
| 120.240.35.178:22222 | ✓ 1037ms | ✓ 1376ms | ✓ 1158ms | ✓ 1336ms | ✓ 1077ms | http |
| 120.240.35.161:22222 | ✓ 1053ms | ✓ 1394ms | ✓ 1147ms | ✓ 1293ms | ✓ 1119ms | http |
| 120.198.141.79:22222 | ✓ 1152ms | ✓ 1483ms | ✓ 1073ms | ✓ 1338ms | ✓ 1095ms | http |
| 185.243.218.43:49153 | ✓ 649ms | 否 | ✓ 1534ms | 否 | ✓ 1602ms | http |
| 120.41.7.123:9091 | ✓ 1800ms | ✓ 1767ms | 否 | ✓ 1965ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1692ms | 否 | ✓ 1147ms | ✓ 1517ms | ✓ 1480ms | http |
| 121.43.196.210:8222 | ✓ 1106ms | ✓ 1269ms | ✓ 1079ms | ✓ 1312ms | ✓ 1068ms | http |
| 121.43.196.213:8222 | ✓ 1124ms | ✓ 1297ms | ✓ 1050ms | ✓ 1338ms | ✓ 1062ms | http |
| 114.55.226.123:10086 | ✓ 1239ms | ✓ 1670ms | ✓ 1189ms | ✓ 1493ms | ✓ 1257ms | http |
| 62.113.119.14:8080 | ✓ 1059ms | ✓ 1627ms | ✓ 1143ms | ✓ 1590ms | ✓ 1156ms | http |
| 115.231.181.40:8128 | ✓ 1064ms | ✓ 1390ms | ✓ 1090ms | ✓ 1754ms | ✓ 1510ms | http |
| 120.92.212.16:7890 | ✓ 1139ms | 否 | 否 | ✓ 1453ms | ✓ 1141ms | http |
| 162.240.154.26:3128 | ✓ 1988ms | ✓ 1684ms | 否 | 否 | ✓ 1668ms | http |
| 109.73.195.10:8888 | ✓ 1176ms | ✓ 1810ms | ✓ 1460ms | 否 | ✓ 1760ms | http |
| 81.70.169.194:80 | ✓ 1152ms | 否 | ✓ 1126ms | ✓ 1482ms | ✓ 1229ms | http |
| 103.215.36.88:16099 | ✓ 1407ms | 否 | ✓ 1421ms | 否 | ✓ 1435ms | http |
| 210.223.44.230:3128 | ✓ 1613ms | ✓ 1136ms | ✓ 1234ms | ✓ 1200ms | ✓ 1230ms | http |
| 59.46.216.131:30001 | ✓ 1157ms | ✓ 1724ms | ✓ 1262ms | ✓ 1524ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1176ms | ✓ 1955ms | 否 | ✓ 1848ms | ✓ 1658ms | http |
| 45.136.198.40:3128 | ✓ 1192ms | ✓ 1994ms | ✓ 1633ms | 否 | ✓ 1642ms | http |
| 83.219.250.8:62920 | ✓ 640ms | 否 | ✓ 1574ms | 否 | ✓ 1342ms | http |
| 117.159.239.44:22222 | ✓ 1048ms | ✓ 1274ms | ✓ 964ms | ✓ 1722ms | ✓ 1563ms | http |
| 113.59.32.141:22222 | ✓ 1284ms | ✓ 1546ms | ✓ 1774ms | ✓ 1474ms | ✓ 1328ms | http |
| 120.240.35.160:22222 | ✓ 1215ms | ✓ 1783ms | ✓ 1941ms | 否 | ✓ 1188ms | http |
| 103.215.36.88:18721 | ✓ 1273ms | 否 | 否 | ✓ 1702ms | ✓ 1285ms | http |
| 58.220.95.8:10174 | ✓ 1347ms | 否 | ✓ 1577ms | ✓ 1538ms | ✓ 1945ms | http |
| 222.184.48.241:22222 | ✓ 1111ms | 否 | 否 | ✓ 1471ms | ✓ 1239ms | http |
| 120.240.35.177:22222 | 否 | 否 | ✓ 1389ms | ✓ 1696ms | ✓ 1198ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1380ms | ✓ 1244ms | ✓ 1556ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1563ms | 否 | ✓ 1402ms | 否 | ✓ 1101ms | http |
| 202.129.206.239:3128 | ✓ 1894ms | 否 | ✓ 1768ms | 否 | ✓ 1874ms | http |
| 121.230.8.60:1080 | ✓ 1288ms | 否 | ✓ 1249ms | ✓ 1609ms | ✓ 1253ms | http |
| 46.249.103.192:443 | ✓ 923ms | 否 | ✓ 1791ms | ✓ 1774ms | 否 | http |
| 172.212.68.37:3128 | ✓ 322ms | 否 | ✓ 1397ms | ✓ 1292ms | ✓ 1080ms | http |
| 47.77.193.180:1080 | ✓ 669ms | ✓ 1626ms | ✓ 412ms | ✓ 1088ms | ✓ 804ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1409ms | ✓ 1383ms | ✓ 1322ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
