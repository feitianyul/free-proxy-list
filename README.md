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

最后更新：2026-03-22 12:17:39 UTC（2026-03-22 20:17:39 UTC+8）

**代理总数：51**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1525ms | 否 | ✓ 992ms | ✓ 1131ms | ✓ 988ms | http |
| 1.231.81.166:3128 | ✓ 1555ms | ✓ 1353ms | ✓ 1474ms | ✓ 1001ms | ✓ 851ms | http |
| 113.160.132.26:8080 | ✓ 1526ms | ✓ 1439ms | 否 | ✓ 1451ms | ✓ 1014ms | http |
| 45.167.124.52:8080 | ✓ 1432ms | 否 | ✓ 842ms | ✓ 1636ms | ✓ 1338ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 542ms | ✓ 1030ms | ✓ 888ms | http |
| 85.198.96.242:3128 | ✓ 688ms | 否 | ✓ 698ms | ✓ 1756ms | ✓ 1716ms | http |
| 167.103.34.108:8800 | ✓ 1140ms | 否 | ✓ 1136ms | ✓ 1349ms | ✓ 1277ms | http |
| 91.238.105.64:2024 | ✓ 1103ms | 否 | ✓ 1610ms | ✓ 1679ms | ✓ 1241ms | http |
| 167.103.31.122:8800 | 否 | 否 | ✓ 1598ms | ✓ 1670ms | ✓ 1576ms | http |
| 142.171.224.229:7890 | ✓ 334ms | ✓ 773ms | ✓ 784ms | ✓ 792ms | ✓ 564ms | http |
| 103.84.95.54:7890 | ✓ 1229ms | 否 | ✓ 1178ms | ✓ 1911ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1803ms | ✓ 1214ms | ✓ 1494ms | ✓ 1445ms | http |
| 181.78.194.249:999 | ✓ 1211ms | ✓ 1971ms | ✓ 1965ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 996ms | ✓ 1258ms | ✓ 1054ms | ✓ 1242ms | ✓ 1238ms | http |
| 147.161.239.240:8800 | ✓ 1228ms | ✓ 1686ms | ✓ 1574ms | 否 | ✓ 1703ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1180ms | ✓ 1135ms | 否 | ✓ 1773ms | http |
| 106.75.15.167:7890 | 否 | ✓ 1171ms | ✓ 954ms | 否 | ✓ 1406ms | http |
| 218.89.134.230:3333 | ✓ 1407ms | 否 | ✓ 1499ms | ✓ 1592ms | ✓ 1429ms | http |
| 34.150.20.6:8888 | ✓ 701ms | ✓ 1175ms | ✓ 1012ms | ✓ 881ms | ✓ 710ms | http |
| 219.117.204.211:7799 | ✓ 1614ms | 否 | ✓ 1915ms | 否 | ✓ 1135ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1333ms | ✓ 909ms | ✓ 1082ms | ✓ 1937ms | http |
| 45.151.183.183:1080 | ✓ 939ms | ✓ 1924ms | 否 | 否 | ✓ 1473ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1128ms | ✓ 1775ms | ✓ 1628ms | 否 | http |
| 150.241.77.172:1080 | ✓ 747ms | 否 | ✓ 981ms | 否 | ✓ 1848ms | http |
| 164.90.155.209:3128 | ✓ 342ms | 否 | ✓ 814ms | ✓ 749ms | ✓ 574ms | http |
| 150.249.255.91:3128 | ✓ 1414ms | 否 | ✓ 583ms | ✓ 853ms | ✓ 687ms | http |
| 47.101.149.27:9010 | ✓ 1378ms | ✓ 1321ms | ✓ 1363ms | ✓ 1438ms | 否 | http |
| 37.187.109.70:10111 | ✓ 629ms | ✓ 1688ms | ✓ 593ms | ✓ 1660ms | ✓ 1343ms | http |
| 185.241.5.57:3128 | ✓ 1120ms | 否 | ✓ 1200ms | 否 | ✓ 1704ms | http |
| 5.102.109.41:999 | ✓ 1414ms | 否 | ✓ 1318ms | ✓ 1211ms | ✓ 1030ms | http |
| 212.192.12.90:6005 | ✓ 1756ms | 否 | ✓ 1994ms | 否 | ✓ 1972ms | http |
| 47.101.159.19:8899 | ✓ 893ms | ✓ 1113ms | ✓ 990ms | ✓ 1151ms | ✓ 882ms | http |
| 194.67.99.223:1080 | ✓ 1132ms | 否 | ✓ 1541ms | ✓ 1923ms | ✓ 1488ms | http |
| 181.78.44.63:999 | ✓ 1729ms | 否 | ✓ 1369ms | ✓ 1760ms | 否 | http |
| 166.88.55.83:7890 | ✓ 655ms | ✓ 1116ms | ✓ 647ms | ✓ 824ms | ✓ 665ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1800ms | ✓ 1972ms | ✓ 1856ms | http |
| 116.80.49.163:3172 | ✓ 1525ms | 否 | ✓ 1522ms | 否 | ✓ 1763ms | http |
| 24.144.86.173:1080 | ✓ 683ms | 否 | ✓ 1275ms | ✓ 838ms | ✓ 658ms | http |
| 185.114.73.2:1080 | ✓ 1018ms | ✓ 1775ms | ✓ 1203ms | 否 | 否 | http |
| 128.199.114.189:9090 | ✓ 1786ms | 否 | ✓ 1590ms | 否 | ✓ 1302ms | http |
| 172.212.68.37:3128 | ✓ 516ms | ✓ 1804ms | ✓ 715ms | ✓ 1620ms | ✓ 1916ms | http |
| 103.39.51.190:8080 | ✓ 1818ms | 否 | 否 | ✓ 1392ms | ✓ 1344ms | http |
| 143.189.3.198:8080 | ✓ 577ms | ✓ 1331ms | ✓ 518ms | ✓ 885ms | ✓ 1648ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 232ms | ✓ 1133ms | ✓ 876ms | http |
| 94.16.114.3:40000 | 否 | 否 | ✓ 1352ms | ✓ 1865ms | ✓ 1835ms | http |
| 45.136.198.40:3128 | ✓ 1226ms | 否 | ✓ 1632ms | 否 | ✓ 1376ms | http |
| 104.129.202.127:10810 | 否 | ✓ 908ms | ✓ 885ms | ✓ 920ms | ✓ 779ms | http |
| 62.113.119.14:8080 | ✓ 1365ms | 否 | ✓ 1350ms | 否 | ✓ 1187ms | http |
| 88.80.150.82:8080 | ✓ 1140ms | 否 | ✓ 1484ms | 否 | ✓ 1826ms | https |
| 45.168.238.193:8443 | ✓ 396ms | 否 | ✓ 191ms | ✓ 1159ms | ✓ 1077ms | http |
| 165.225.113.220:10003 | ✓ 737ms | 否 | ✓ 732ms | ✓ 1058ms | 否 | http |

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
