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

最后更新：2026-03-03 18:41:30 UTC（2026-03-04 02:41:30 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 540ms | 否 | ✓ 958ms | ✓ 1074ms | ✓ 801ms | http |
| 166.0.192.117:8888 | 否 | ✓ 1318ms | ✓ 1876ms | ✓ 1435ms | ✓ 970ms | http |
| 3.213.157.4:3128 | ✓ 358ms | 否 | ✓ 722ms | ✓ 1837ms | ✓ 1363ms | http |
| 217.76.245.80:999 | ✓ 730ms | ✓ 1547ms | ✓ 1070ms | ✓ 1295ms | ✓ 1266ms | http |
| 186.148.180.46:999 | ✓ 1082ms | ✓ 1788ms | ✓ 1200ms | ✓ 1620ms | ✓ 1375ms | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 1188ms | ✓ 1231ms | ✓ 980ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1362ms | ✓ 1460ms | ✓ 1117ms | http |
| 74.208.234.198:443 | ✓ 820ms | 否 | 否 | ✓ 1830ms | ✓ 1151ms | http |
| 46.249.103.192:443 | ✓ 1340ms | 否 | ✓ 1641ms | ✓ 1924ms | 否 | http |
| 35.234.17.221:8080 | ✓ 1007ms | 否 | ✓ 1053ms | ✓ 1257ms | 否 | http |
| 35.225.22.61:80 | ✓ 806ms | ✓ 1144ms | ✓ 939ms | ✓ 1031ms | ✓ 731ms | http |
| 147.45.251.242:8888 | ✓ 1908ms | ✓ 1971ms | ✓ 1362ms | 否 | 否 | http |
| 91.193.240.157:9877 | ✓ 866ms | ✓ 1820ms | ✓ 743ms | 否 | ✓ 1347ms | http |
| 103.189.139.254:8080 | 否 | 否 | ✓ 1856ms | ✓ 1760ms | ✓ 1676ms | http |
| 103.84.95.54:7890 | ✓ 839ms | 否 | 否 | ✓ 1254ms | ✓ 911ms | http |
| 207.254.71.62:8088 | ✓ 603ms | ✓ 1744ms | 否 | ✓ 1900ms | ✓ 1474ms | http |
| 172.212.68.37:3128 | ✓ 778ms | 否 | ✓ 728ms | ✓ 1768ms | 否 | http |
| 101.43.255.96:80 | ✓ 1613ms | ✓ 1562ms | ✓ 1151ms | ✓ 1551ms | ✓ 1128ms | http |
| 115.231.181.40:8128 | ✓ 1612ms | 否 | ✓ 1680ms | 否 | ✓ 1107ms | http |
| 178.156.224.42:3128 | ✓ 991ms | ✓ 1977ms | 否 | 否 | ✓ 1647ms | http |
| 34.101.184.164:3128 | ✓ 1207ms | 否 | ✓ 1001ms | ✓ 1449ms | ✓ 1150ms | http |
| 210.223.44.230:3128 | ✓ 924ms | 否 | ✓ 777ms | 否 | ✓ 862ms | http |
| 152.70.137.18:8888 | ✓ 1391ms | ✓ 1319ms | 否 | 否 | ✓ 1687ms | http |
| 5.75.196.26:40000 | 否 | ✓ 1617ms | ✓ 1541ms | ✓ 1141ms | ✓ 1556ms | http |
| 81.70.169.194:80 | ✓ 1133ms | ✓ 1528ms | ✓ 1162ms | ✓ 1438ms | ✓ 1196ms | http |
| 24.199.124.152:3128 | ✓ 1077ms | ✓ 1263ms | ✓ 1332ms | ✓ 1000ms | ✓ 752ms | http |
| 116.105.21.153:9003 | ✓ 1360ms | 否 | ✓ 1366ms | ✓ 1819ms | ✓ 1658ms | http |
| 159.89.31.62:8080 | ✓ 459ms | 否 | ✓ 432ms | ✓ 1727ms | ✓ 1567ms | http |
| 165.22.240.238:3128 | 否 | 否 | ✓ 1458ms | ✓ 1283ms | ✓ 1034ms | http |
| 154.26.133.70:3128 | 否 | 否 | ✓ 1783ms | ✓ 1847ms | ✓ 1701ms | http |
| 102.217.69.117:3128 | ✓ 1551ms | 否 | ✓ 1823ms | ✓ 1899ms | 否 | http |
| 144.91.81.25:3128 | ✓ 1350ms | ✓ 1849ms | 否 | 否 | ✓ 1694ms | http |
| 121.230.8.251:1080 | ✓ 1273ms | ✓ 1498ms | ✓ 1383ms | ✓ 1985ms | 否 | http |
| 47.95.231.180:8084 | ✓ 1070ms | ✓ 1454ms | ✓ 1043ms | ✓ 1406ms | ✓ 1096ms | http |
| 61.72.221.234:3128 | 否 | ✓ 1650ms | ✓ 1602ms | ✓ 1477ms | 否 | http |
| 90.84.188.97:8000 | ✓ 659ms | ✓ 1628ms | ✓ 1880ms | 否 | 否 | http |
| 101.32.244.83:8080 | ✓ 1571ms | ✓ 1992ms | ✓ 1143ms | ✓ 1720ms | ✓ 1593ms | http |
| 121.43.196.213:8222 | ✓ 1108ms | ✓ 1232ms | ✓ 1068ms | ✓ 1380ms | ✓ 1031ms | http |
| 121.43.196.210:8222 | ✓ 1102ms | ✓ 1309ms | ✓ 996ms | ✓ 1334ms | ✓ 1097ms | http |
| 114.55.226.123:10086 | ✓ 1204ms | ✓ 1616ms | ✓ 1149ms | ✓ 1520ms | ✓ 1179ms | http |
| 62.113.119.14:8080 | ✓ 1282ms | 否 | ✓ 1196ms | ✓ 1610ms | ✓ 1333ms | http |
| 5.101.0.233:3128 | ✓ 1782ms | ✓ 1439ms | ✓ 1225ms | ✓ 1951ms | ✓ 1690ms | http |
| 162.240.154.26:3128 | ✓ 1160ms | 否 | ✓ 1399ms | ✓ 1565ms | ✓ 1249ms | http |
| 45.136.198.40:3128 | ✓ 660ms | ✓ 1838ms | ✓ 1585ms | 否 | ✓ 1600ms | http |
| 111.79.111.126:3128 | ✓ 1429ms | 否 | 否 | ✓ 1443ms | ✓ 1675ms | http |
| 120.92.212.16:7890 | ✓ 1373ms | ✓ 1393ms | 否 | ✓ 1433ms | ✓ 1108ms | http |
| 121.230.8.136:1080 | ✓ 1373ms | ✓ 1570ms | ✓ 1427ms | 否 | 否 | http |
| 113.255.59.226:8080 | ✓ 1432ms | 否 | ✓ 1541ms | ✓ 1396ms | ✓ 1391ms | http |
| 2.56.178.131:443 | ✓ 1280ms | 否 | ✓ 1206ms | 否 | ✓ 1656ms | http |
| 121.230.9.168:1080 | 否 | 否 | ✓ 1238ms | ✓ 1744ms | ✓ 1655ms | http |
| 103.215.36.88:11929 | 否 | ✓ 1464ms | ✓ 1314ms | ✓ 1656ms | ✓ 1262ms | http |
| 94.158.49.82:3128 | 否 | 否 | ✓ 1351ms | ✓ 1962ms | ✓ 1602ms | http |
| 180.130.80.196:9003 | ✓ 1491ms | 否 | ✓ 1635ms | ✓ 1590ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1201ms | ✓ 1843ms | ✓ 1804ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1913ms | 否 | 否 | ✓ 1941ms | ✓ 1981ms | http |
| 103.215.36.88:16851 | ✓ 1227ms | ✓ 1576ms | ✓ 1284ms | ✓ 1485ms | ✓ 1727ms | http |
| 223.16.170.103:80 | ✓ 1149ms | 否 | ✓ 1268ms | ✓ 1335ms | ✓ 1379ms | http |
| 37.187.109.70:10111 | ✓ 1527ms | 否 | ✓ 971ms | ✓ 1640ms | 否 | http |
| 192.166.82.55:1080 | ✓ 1858ms | 否 | ✓ 868ms | ✓ 1996ms | ✓ 1408ms | http |
| 106.14.203.63:3333 | ✓ 1078ms | ✓ 1275ms | ✓ 1050ms | ✓ 1334ms | ✓ 1062ms | http |

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
