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

最后更新：2026-04-04 08:10:32 UTC（2026-04-04 16:10:32 UTC+8）

**代理总数：58**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 924ms | 否 | ✓ 967ms | ✓ 1168ms | ✓ 949ms | http |
| 147.161.210.140:8800 | ✓ 996ms | 否 | ✓ 1104ms | ✓ 1140ms | ✓ 1499ms | http |
| 95.213.217.168:52004 | ✓ 1203ms | 否 | ✓ 670ms | ✓ 1883ms | ✓ 1795ms | http |
| 113.160.132.26:8080 | ✓ 1533ms | ✓ 1528ms | ✓ 1085ms | ✓ 1460ms | ✓ 1133ms | http |
| 167.103.115.102:8800 | ✓ 1516ms | 否 | ✓ 1151ms | ✓ 1337ms | ✓ 1473ms | http |
| 167.103.34.108:8800 | ✓ 1552ms | 否 | ✓ 1426ms | ✓ 1459ms | ✓ 1269ms | http |
| 159.223.71.162:8080 | ✓ 1516ms | 否 | 否 | ✓ 1268ms | ✓ 1357ms | http |
| 159.223.71.162:443 | ✓ 1517ms | 否 | 否 | ✓ 1274ms | ✓ 1368ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1367ms | ✓ 1373ms | 否 | ✓ 1099ms | http |
| 133.242.138.34:8100 | ✓ 772ms | ✓ 1721ms | ✓ 951ms | 否 | ✓ 1990ms | http |
| 111.227.254.12:22222 | ✓ 1146ms | ✓ 1583ms | 否 | 否 | ✓ 1624ms | http |
| 111.227.254.9:22222 | 否 | 否 | ✓ 1198ms | ✓ 1568ms | ✓ 1281ms | http |
| 45.167.124.52:8080 | ✓ 482ms | ✓ 1460ms | ✓ 691ms | ✓ 1543ms | ✓ 1258ms | http |
| 218.108.131.186:17890 | ✓ 802ms | ✓ 1061ms | ✓ 1285ms | ✓ 1068ms | ✓ 842ms | http |
| 1.231.81.166:3128 | ✓ 918ms | ✓ 1685ms | ✓ 1649ms | ✓ 1131ms | ✓ 922ms | http |
| 167.103.144.127:8800 | ✓ 1215ms | ✓ 1978ms | ✓ 1256ms | ✓ 1798ms | ✓ 1317ms | http |
| 167.103.31.122:8800 | ✓ 1266ms | 否 | ✓ 1976ms | 否 | ✓ 1829ms | http |
| 35.225.22.61:80 | ✓ 972ms | 否 | 否 | ✓ 1060ms | ✓ 765ms | http |
| 101.43.127.100:8877 | ✓ 1475ms | ✓ 1241ms | 否 | ✓ 1631ms | ✓ 1524ms | http |
| 45.167.125.21:999 | ✓ 527ms | ✓ 1438ms | ✓ 1337ms | ✓ 1660ms | ✓ 1646ms | http |
| 181.78.44.63:999 | ✓ 840ms | ✓ 1366ms | 否 | ✓ 1604ms | ✓ 1344ms | http |
| 177.234.217.88:999 | ✓ 1151ms | 否 | ✓ 1832ms | 否 | ✓ 1582ms | http |
| 82.114.228.67:1080 | ✓ 960ms | 否 | ✓ 841ms | 否 | ✓ 1168ms | http |
| 64.227.76.27:1080 | ✓ 416ms | 否 | ✓ 1254ms | 否 | ✓ 1412ms | http |
| 80.250.165.242:3128 | ✓ 991ms | ✓ 1873ms | ✓ 1394ms | 否 | ✓ 1702ms | http |
| 1.225.116.115:1080 | ✓ 1307ms | 否 | ✓ 1029ms | ✓ 1135ms | ✓ 909ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 893ms | ✓ 1504ms | ✓ 1125ms | http |
| 45.140.147.82:1081 | ✓ 948ms | 否 | ✓ 578ms | 否 | ✓ 1029ms | http |
| 61.76.95.217:40088 | ✓ 1340ms | 否 | ✓ 1232ms | ✓ 1507ms | ✓ 1221ms | http |
| 118.113.246.49:1080 | ✓ 1421ms | ✓ 1959ms | 否 | ✓ 1693ms | ✓ 1569ms | http |
| 57.128.188.167:9163 | ✓ 1820ms | 否 | ✓ 1620ms | 否 | ✓ 1727ms | http |
| 101.132.61.121:8888 | ✓ 1535ms | ✓ 1426ms | ✓ 1561ms | ✓ 1679ms | ✓ 1487ms | http |
| 103.113.70.189:1081 | ✓ 786ms | 否 | ✓ 68ms | ✓ 1194ms | ✓ 862ms | http |
| 65.21.90.27:1019 | ✓ 1903ms | 否 | ✓ 1894ms | 否 | ✓ 1968ms | http |
| 45.125.67.37:443 | ✓ 1055ms | 否 | ✓ 1252ms | ✓ 1459ms | ✓ 1314ms | http |
| 93.77.179.177:8888 | ✓ 580ms | ✓ 1708ms | ✓ 767ms | ✓ 1753ms | ✓ 1350ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1429ms | ✓ 1162ms | 否 | ✓ 1179ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1474ms | ✓ 1281ms | ✓ 1450ms | ✓ 1182ms | http |
| 210.223.44.230:3128 | ✓ 964ms | ✓ 1262ms | ✓ 1166ms | ✓ 1244ms | ✓ 972ms | http |
| 209.38.154.7:1080 | 否 | ✓ 895ms | ✓ 1465ms | ✓ 912ms | ✓ 699ms | http |
| 43.167.237.94:3128 | ✓ 1653ms | 否 | ✓ 1259ms | ✓ 1479ms | ✓ 1575ms | http |
| 59.46.216.131:30001 | ✓ 1545ms | ✓ 1531ms | 否 | 否 | ✓ 1222ms | http |
| 194.87.215.166:9443 | ✓ 1125ms | 否 | ✓ 992ms | 否 | ✓ 1856ms | http |
| 217.217.249.160:8080 | ✓ 1657ms | 否 | ✓ 1755ms | 否 | ✓ 1390ms | http |
| 5.104.87.17:8051 | ✓ 1968ms | 否 | 否 | ✓ 1685ms | ✓ 1940ms | http |
| 24.144.86.173:1080 | ✓ 741ms | ✓ 991ms | ✓ 949ms | ✓ 918ms | ✓ 704ms | http |
| 180.103.19.220:1080 | 否 | ✓ 1499ms | ✓ 1185ms | 否 | ✓ 1360ms | http |
| 114.237.77.244:1080 | ✓ 1527ms | ✓ 1430ms | ✓ 1149ms | ✓ 1555ms | 否 | http |
| 195.123.209.48:3128 | ✓ 807ms | ✓ 1397ms | ✓ 1346ms | 否 | 否 | http |
| 38.145.218.238:8447 | ✓ 975ms | ✓ 1548ms | ✓ 742ms | ✓ 895ms | ✓ 1389ms | http |
| 8.219.97.248:80 | ✓ 1582ms | 否 | ✓ 1192ms | ✓ 1766ms | ✓ 1516ms | http |
| 45.129.141.143:3128 | ✓ 1045ms | ✓ 1668ms | 否 | ✓ 1959ms | ✓ 1398ms | http |
| 212.58.132.5:8888 | ✓ 1156ms | 否 | ✓ 1889ms | ✓ 1947ms | ✓ 1268ms | http |
| 207.244.244.178:3128 | ✓ 418ms | ✓ 1118ms | ✓ 1370ms | ✓ 1230ms | ✓ 924ms | http |
| 109.234.38.35:3128 | ✓ 1479ms | ✓ 1524ms | ✓ 1560ms | ✓ 1942ms | ✓ 1647ms | http |
| 185.114.73.2:1080 | ✓ 1366ms | ✓ 1729ms | ✓ 1586ms | 否 | ✓ 1730ms | http |
| 208.87.243.199:7878 | ✓ 545ms | ✓ 1344ms | ✓ 910ms | ✓ 1937ms | ✓ 725ms | http |
| 92.119.127.211:6005 | 否 | ✓ 1399ms | ✓ 1725ms | ✓ 1748ms | ✓ 1728ms | http |

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
