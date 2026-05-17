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

最后更新：2026-05-17 08:25:09 UTC（2026-05-17 16:25:09 UTC+8）

**代理总数：46**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 51.161.50.166:3128 | ✓ 281ms | 否 | ✓ 1814ms | 否 | ✓ 1967ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1578ms | ✓ 1512ms | ✓ 1590ms | ✓ 1336ms | http |
| 170.106.136.181:31002 | ✓ 872ms | ✓ 1047ms | ✓ 1247ms | ✓ 1225ms | ✓ 1688ms | http |
| 185.200.188.234:10001 | ✓ 1109ms | 否 | ✓ 1324ms | 否 | ✓ 1697ms | http |
| 185.40.77.94:1080 | ✓ 1428ms | 否 | ✓ 1482ms | 否 | ✓ 1918ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1065ms | ✓ 1362ms | ✓ 1326ms | http |
| 218.108.131.186:17890 | ✓ 1051ms | ✓ 1273ms | ✓ 1080ms | ✓ 1265ms | ✓ 1063ms | http |
| 65.109.125.111:8443 | ✓ 603ms | 否 | ✓ 1328ms | 否 | ✓ 1946ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1554ms | ✓ 1369ms | ✓ 1106ms | http |
| 8.154.21.175:3128 | ✓ 1362ms | ✓ 1277ms | ✓ 1587ms | 否 | 否 | http |
| 148.230.4.241:999 | ✓ 828ms | ✓ 1649ms | ✓ 745ms | ✓ 1819ms | ✓ 1311ms | http |
| 5.75.139.30:1081 | ✓ 561ms | ✓ 1912ms | ✓ 1274ms | ✓ 1884ms | ✓ 1619ms | http |
| 114.214.165.78:10810 | ✓ 1612ms | 否 | ✓ 1762ms | ✓ 1682ms | ✓ 1688ms | http |
| 158.255.212.55:9005 | ✓ 851ms | 否 | ✓ 1365ms | ✓ 1664ms | 否 | http |
| 158.255.212.55:7497 | ✓ 852ms | 否 | ✓ 1363ms | ✓ 1657ms | 否 | http |
| 158.255.212.55:9480 | ✓ 848ms | 否 | ✓ 1368ms | ✓ 1653ms | 否 | http |
| 147.45.78.89:1080 | ✓ 1775ms | 否 | ✓ 1533ms | ✓ 1836ms | ✓ 1355ms | http |
| 43.156.90.221:10808 | ✓ 1121ms | 否 | ✓ 1481ms | ✓ 1788ms | 否 | http |
| 113.45.216.128:55555 | 否 | 否 | ✓ 1373ms | ✓ 1648ms | ✓ 1343ms | http |
| 170.78.208.1:999 | ✓ 1074ms | 否 | ✓ 997ms | ✓ 1252ms | ✓ 1460ms | http |
| 91.242.229.129:8092 | 否 | ✓ 1502ms | ✓ 1981ms | ✓ 1523ms | ✓ 1149ms | http |
| 86.104.74.110:1081 | ✓ 1139ms | 否 | ✓ 847ms | 否 | ✓ 1820ms | http |
| 84.47.150.125:1080 | ✓ 1027ms | 否 | ✓ 1687ms | 否 | ✓ 1861ms | http |
| 166.88.55.83:7890 | ✓ 1216ms | ✓ 1396ms | ✓ 856ms | ✓ 1106ms | ✓ 807ms | http |
| 137.184.0.30:3128 | ✓ 1560ms | 否 | ✓ 1099ms | ✓ 946ms | 否 | http |
| 3.15.187.17:1080 | ✓ 316ms | 否 | ✓ 771ms | 否 | ✓ 906ms | http |
| 45.125.67.37:8443 | ✓ 1695ms | 否 | 否 | ✓ 1320ms | ✓ 1413ms | http |
| 129.80.238.83:444 | ✓ 555ms | ✓ 987ms | ✓ 542ms | 否 | 否 | http |
| 158.255.212.55:3256 | ✓ 1828ms | 否 | ✓ 1724ms | ✓ 1838ms | 否 | http |
| 57.129.144.178:40000 | ✓ 949ms | 否 | ✓ 1532ms | ✓ 1637ms | ✓ 1355ms | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1896ms | ✓ 1458ms | ✓ 1178ms | http |
| 159.223.41.216:9090 | ✓ 922ms | 否 | ✓ 914ms | ✓ 1282ms | ✓ 1070ms | http |
| 103.21.220.141:3128 | ✓ 1303ms | 否 | 否 | ✓ 1231ms | ✓ 1000ms | http |
| 210.223.44.230:3128 | ✓ 1990ms | ✓ 1427ms | ✓ 865ms | 否 | ✓ 994ms | http |
| 3.101.133.120:80 | ✓ 906ms | 否 | ✓ 1769ms | ✓ 1312ms | ✓ 1320ms | http |
| 185.71.196.92:1080 | ✓ 851ms | 否 | ✓ 996ms | 否 | ✓ 1809ms | http |
| 217.182.195.221:30001 | ✓ 1691ms | 否 | ✓ 798ms | 否 | ✓ 1949ms | http |
| 212.58.132.5:8888 | ✓ 1316ms | 否 | ✓ 1042ms | ✓ 1452ms | ✓ 1139ms | http |
| 103.147.152.12:1080 | ✓ 910ms | 否 | ✓ 1133ms | 否 | ✓ 1294ms | http |
| 45.129.141.143:3128 | ✓ 1193ms | ✓ 1813ms | ✓ 1770ms | ✓ 1757ms | ✓ 1648ms | http |
| 193.160.209.58:1080 | ✓ 1131ms | 否 | ✓ 1386ms | ✓ 1911ms | ✓ 1682ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 1140ms | ✓ 1861ms | ✓ 1270ms | http |
| 42.200.76.16:3888 | ✓ 1121ms | 否 | ✓ 967ms | ✓ 1109ms | ✓ 1867ms | http |
| 61.52.131.172:8443 | ✓ 1070ms | ✓ 1430ms | ✓ 1131ms | ✓ 1415ms | ✓ 1128ms | http |
| 207.254.71.62:8088 | ✓ 1303ms | 否 | ✓ 1607ms | ✓ 1756ms | ✓ 1843ms | http |
| 59.46.216.131:30001 | ✓ 1205ms | 否 | ✓ 1308ms | ✓ 1679ms | 否 | http |

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
