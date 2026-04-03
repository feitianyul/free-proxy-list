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

最后更新：2026-04-03 20:39:41 UTC（2026-04-04 04:39:41 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 916ms | ✓ 1238ms | ✓ 1044ms | ✓ 1152ms | ✓ 1054ms | http |
| 1.231.81.166:3128 | ✓ 1143ms | ✓ 1126ms | ✓ 1895ms | ✓ 1346ms | ✓ 1023ms | http |
| 95.213.217.168:52004 | ✓ 894ms | ✓ 1681ms | ✓ 910ms | 否 | ✓ 1641ms | http |
| 111.227.254.9:22222 | ✓ 1095ms | ✓ 1379ms | ✓ 1059ms | ✓ 1445ms | ✓ 1084ms | http |
| 111.227.254.12:22222 | ✓ 1146ms | ✓ 1453ms | ✓ 1181ms | ✓ 1496ms | ✓ 1218ms | http |
| 167.103.115.102:8800 | ✓ 1727ms | 否 | ✓ 1069ms | ✓ 1280ms | ✓ 1393ms | http |
| 203.80.138.81:50000 | 否 | ✓ 1198ms | ✓ 1011ms | ✓ 1031ms | 否 | http |
| 139.159.99.242:8080 | ✓ 892ms | 否 | 否 | ✓ 1924ms | ✓ 1371ms | http |
| 113.160.132.26:8080 | ✓ 1546ms | ✓ 1553ms | ✓ 1425ms | ✓ 1867ms | ✓ 1463ms | http |
| 167.103.34.108:8800 | ✓ 1779ms | 否 | ✓ 1633ms | ✓ 1470ms | ✓ 1525ms | http |
| 218.108.131.186:17890 | ✓ 1171ms | ✓ 1134ms | ✓ 1766ms | ✓ 1227ms | ✓ 1704ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1901ms | ✓ 1368ms | 否 | ✓ 1614ms | http |
| 64.227.76.27:1080 | ✓ 1172ms | ✓ 1933ms | ✓ 1391ms | 否 | 否 | http |
| 45.149.92.147:5001 | 否 | 否 | ✓ 957ms | ✓ 1565ms | ✓ 1312ms | http |
| 167.103.144.127:8800 | ✓ 1346ms | 否 | ✓ 1555ms | 否 | ✓ 1601ms | http |
| 167.103.31.122:8800 | ✓ 1585ms | 否 | ✓ 1335ms | ✓ 1647ms | ✓ 1562ms | http |
| 101.43.127.100:8877 | ✓ 983ms | ✓ 1243ms | ✓ 1120ms | ✓ 1199ms | ✓ 1302ms | http |
| 147.161.239.240:8800 | ✓ 1119ms | ✓ 1716ms | ✓ 1115ms | ✓ 1611ms | ✓ 1485ms | http |
| 120.92.212.16:7890 | ✓ 1586ms | 否 | ✓ 1364ms | 否 | ✓ 1046ms | http |
| 208.87.243.199:7878 | ✓ 1544ms | ✓ 1554ms | ✓ 1466ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1027ms | 否 | 否 | ✓ 1600ms | ✓ 1534ms | http |
| 59.46.216.131:30001 | ✓ 1008ms | ✓ 1461ms | 否 | 否 | ✓ 1143ms | http |
| 46.39.105.157:8080 | ✓ 806ms | ✓ 1900ms | ✓ 1876ms | 否 | 否 | http |
| 47.238.220.4:8888 | 否 | ✓ 1846ms | ✓ 950ms | ✓ 1536ms | 否 | http |
| 192.71.213.85:9090 | ✓ 1361ms | 否 | ✓ 1645ms | ✓ 1937ms | 否 | http |
| 159.223.71.162:443 | ✓ 838ms | 否 | ✓ 1486ms | ✓ 1129ms | ✓ 944ms | http |
| 159.223.71.162:8080 | ✓ 832ms | 否 | ✓ 1480ms | ✓ 1143ms | ✓ 953ms | http |
| 35.225.22.61:80 | ✓ 300ms | ✓ 1233ms | ✓ 819ms | ✓ 1038ms | ✓ 799ms | http |
| 34.101.184.164:3128 | ✓ 1613ms | 否 | ✓ 1427ms | ✓ 1397ms | ✓ 1096ms | http |
| 45.12.151.226:2829 | ✓ 1268ms | 否 | ✓ 1069ms | 否 | ✓ 1860ms | http |
| 212.58.132.5:8888 | ✓ 1484ms | 否 | ✓ 1235ms | ✓ 1526ms | ✓ 1304ms | http |
| 115.231.181.40:8128 | ✓ 1199ms | ✓ 1173ms | ✓ 1775ms | 否 | 否 | http |
| 72.11.151.159:6005 | ✓ 555ms | ✓ 1249ms | ✓ 1415ms | ✓ 1601ms | ✓ 1196ms | http |
| 45.136.198.40:3128 | ✓ 667ms | ✓ 1550ms | ✓ 1713ms | ✓ 1967ms | ✓ 1763ms | http |
| 86.53.183.16:1080 | ✓ 1573ms | 否 | ✓ 919ms | 否 | ✓ 1656ms | http |
| 5.104.87.17:8051 | ✓ 1005ms | 否 | ✓ 905ms | ✓ 1102ms | 否 | http |
| 180.125.216.109:8118 | 否 | ✓ 1308ms | ✓ 1064ms | 否 | ✓ 989ms | http |
| 172.245.67.195:7890 | ✓ 583ms | 否 | ✓ 1891ms | 否 | ✓ 1145ms | http |
| 150.249.255.91:3128 | ✓ 671ms | 否 | ✓ 584ms | ✓ 968ms | ✓ 1301ms | http |
| 38.145.220.60:8447 | ✓ 1551ms | 否 | ✓ 1637ms | ✓ 821ms | ✓ 1998ms | http |
| 103.156.233.137:8080 | 否 | 否 | ✓ 1690ms | ✓ 1469ms | ✓ 1813ms | http |
| 65.108.203.36:18080 | ✓ 845ms | ✓ 1951ms | 否 | 否 | ✓ 1933ms | http |
| 210.223.44.230:3128 | ✓ 688ms | 否 | 否 | ✓ 1041ms | ✓ 787ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1194ms | 否 | ✓ 1487ms | ✓ 1137ms | http |
| 8.219.97.248:80 | ✓ 1710ms | 否 | 否 | ✓ 1944ms | ✓ 1733ms | http |
| 18.201.114.187:50537 | ✓ 1282ms | 否 | ✓ 1403ms | 否 | ✓ 1964ms | http |
| 3.79.194.222:23441 | ✓ 1256ms | 否 | ✓ 1891ms | 否 | ✓ 1840ms | http |
| 195.123.209.48:3128 | ✓ 1220ms | ✓ 1769ms | ✓ 1394ms | 否 | ✓ 1835ms | http |
| 103.113.70.189:1081 | ✓ 1336ms | ✓ 1083ms | ✓ 229ms | ✓ 1142ms | ✓ 1012ms | http |
| 147.45.186.28:3128 | ✓ 1559ms | ✓ 1998ms | ✓ 858ms | ✓ 1901ms | ✓ 1829ms | http |
| 185.41.152.110:3128 | ✓ 1616ms | 否 | ✓ 1937ms | ✓ 1947ms | 否 | http |
| 178.156.224.42:3128 | ✓ 1362ms | ✓ 1876ms | ✓ 1794ms | 否 | 否 | http |
| 18.100.126.55:3480 | ✓ 1193ms | 否 | ✓ 974ms | 否 | ✓ 1529ms | http |

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
