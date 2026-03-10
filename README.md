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

最后更新：2026-03-10 05:38:06 UTC（2026-03-10 13:38:06 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 117ms | ✓ 1947ms | ✓ 1074ms | ✓ 840ms | ✓ 528ms | http |
| 1.231.81.166:3128 | ✓ 1464ms | ✓ 1286ms | ✓ 1307ms | ✓ 918ms | ✓ 742ms | http |
| 154.3.236.202:3128 | ✓ 572ms | 否 | ✓ 1332ms | ✓ 1916ms | ✓ 1143ms | http |
| 202.155.12.161:443 | ✓ 1510ms | 否 | 否 | ✓ 1747ms | ✓ 1433ms | http |
| 91.107.141.42:8081 | ✓ 1209ms | 否 | ✓ 1359ms | 否 | ✓ 1698ms | http |
| 101.47.73.135:3128 | ✓ 1018ms | 否 | ✓ 969ms | 否 | ✓ 1360ms | http |
| 162.240.154.26:3128 | ✓ 826ms | 否 | ✓ 1345ms | ✓ 1646ms | ✓ 1526ms | http |
| 115.231.181.40:8128 | ✓ 905ms | 否 | ✓ 896ms | ✓ 1759ms | 否 | http |
| 116.80.49.166:3172 | ✓ 1827ms | 否 | 否 | ✓ 1785ms | ✓ 1635ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1223ms | ✓ 977ms | ✓ 1214ms | 否 | http |
| 35.225.22.61:80 | ✓ 886ms | 否 | ✓ 437ms | ✓ 1339ms | ✓ 970ms | http |
| 81.70.169.194:80 | ✓ 1090ms | ✓ 1755ms | ✓ 1926ms | ✓ 1317ms | ✓ 1064ms | http |
| 101.43.255.96:80 | ✓ 1021ms | ✓ 1982ms | ✓ 1171ms | 否 | 否 | http |
| 39.104.201.40:7890 | ✓ 1552ms | ✓ 1748ms | 否 | 否 | ✓ 1604ms | http |
| 190.9.109.198:999 | ✓ 1077ms | 否 | ✓ 1346ms | ✓ 1651ms | ✓ 1216ms | http |
| 121.237.181.137:8888 | ✓ 948ms | 否 | ✓ 1441ms | ✓ 1101ms | ✓ 909ms | http |
| 193.168.173.136:443 | ✓ 1048ms | 否 | ✓ 1327ms | 否 | ✓ 1565ms | http |
| 125.128.12.14:3128 | ✓ 1106ms | ✓ 1379ms | ✓ 604ms | ✓ 943ms | ✓ 741ms | http |
| 194.213.18.200:443 | ✓ 1308ms | 否 | ✓ 771ms | 否 | ✓ 1295ms | http |
| 116.236.189.93:29999 | ✓ 1246ms | 否 | ✓ 756ms | 否 | ✓ 770ms | http |
| 113.177.131.2:3128 | ✓ 1412ms | ✓ 1691ms | ✓ 799ms | ✓ 1523ms | ✓ 1277ms | http |
| 5.101.0.233:3128 | ✓ 1113ms | 否 | ✓ 1180ms | ✓ 1800ms | ✓ 1548ms | http |
| 116.80.49.159:3172 | ✓ 1584ms | 否 | ✓ 1467ms | ✓ 1876ms | 否 | http |
| 45.136.130.223:8443 | ✓ 535ms | ✓ 606ms | ✓ 87ms | ✓ 656ms | ✓ 493ms | http |
| 14.225.222.213:7890 | 否 | 否 | ✓ 1157ms | ✓ 1164ms | ✓ 1528ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1230ms | 否 | ✓ 1883ms | ✓ 1322ms | http |
| 165.227.5.10:8888 | ✓ 111ms | 否 | ✓ 577ms | 否 | ✓ 476ms | http |
| 175.215.60.67:3128 | ✓ 1404ms | ✓ 1501ms | ✓ 841ms | ✓ 991ms | 否 | http |
| 95.213.247.248:1080 | ✓ 1838ms | 否 | ✓ 1702ms | 否 | ✓ 1810ms | http |
| 211.171.114.154:3128 | ✓ 1291ms | ✓ 1158ms | ✓ 1209ms | 否 | ✓ 922ms | http |
| 46.183.25.8:443 | ✓ 1881ms | 否 | ✓ 1573ms | ✓ 1806ms | 否 | http |
| 152.70.98.46:8888 | ✓ 651ms | 否 | ✓ 1338ms | ✓ 1260ms | ✓ 1509ms | http |
| 116.80.49.165:3172 | ✓ 1459ms | 否 | ✓ 1538ms | 否 | ✓ 1737ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1352ms | ✓ 772ms | ✓ 1286ms | ✓ 968ms | http |
| 121.230.9.54:1080 | ✓ 1063ms | ✓ 1281ms | 否 | 否 | ✓ 1172ms | http |
| 47.94.105.38:8124 | ✓ 1391ms | 否 | ✓ 1223ms | 否 | ✓ 1617ms | http |
| 45.136.198.40:3128 | ✓ 1395ms | 否 | ✓ 1957ms | 否 | ✓ 1759ms | http |
| 116.80.49.163:3172 | 否 | 否 | ✓ 1989ms | ✓ 1816ms | ✓ 1606ms | http |
| 45.133.251.1:1080 | ✓ 1257ms | 否 | ✓ 715ms | 否 | ✓ 1550ms | http |
| 59.46.216.131:30001 | ✓ 1007ms | 否 | ✓ 1099ms | 否 | ✓ 1004ms | http |
| 45.129.141.143:3128 | ✓ 1127ms | ✓ 1909ms | ✓ 1684ms | 否 | ✓ 1930ms | http |
| 107.178.115.140:3128 | 否 | 否 | ✓ 1657ms | ✓ 1488ms | ✓ 1010ms | http |
| 106.14.203.63:3333 | ✓ 1571ms | ✓ 1051ms | ✓ 884ms | ✓ 1073ms | ✓ 1629ms | http |
| 14.225.212.37:7890 | ✓ 1785ms | 否 | ✓ 1948ms | ✓ 1723ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1743ms | 否 | 否 | ✓ 1436ms | ✓ 1325ms | http |
| 150.107.140.238:3128 | ✓ 1536ms | 否 | ✓ 1075ms | 否 | ✓ 1994ms | http |
| 86.53.183.16:1080 | ✓ 799ms | 否 | ✓ 1344ms | 否 | ✓ 1488ms | http |
| 103.166.185.54:3128 | 否 | ✓ 1990ms | ✓ 1149ms | ✓ 1472ms | ✓ 1269ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1819ms | ✓ 1431ms | ✓ 1410ms | http |
| 116.80.49.168:3172 | ✓ 1834ms | 否 | ✓ 1479ms | ✓ 1788ms | ✓ 1641ms | http |
| 116.80.49.169:3172 | ✓ 1834ms | 否 | ✓ 1478ms | ✓ 1786ms | 否 | http |
| 61.52.131.172:8443 | ✓ 855ms | ✓ 1076ms | ✓ 923ms | ✓ 1196ms | ✓ 934ms | http |
| 178.236.245.59:3128 | ✓ 1434ms | 否 | ✓ 1432ms | ✓ 1883ms | ✓ 1439ms | http |
| 178.236.245.17:3128 | ✓ 1417ms | 否 | ✓ 1430ms | ✓ 1861ms | ✓ 1487ms | http |
| 185.191.236.162:3128 | ✓ 811ms | 否 | ✓ 1544ms | ✓ 1768ms | ✓ 1321ms | http |
| 95.3.9.78:3128 | ✓ 1187ms | 否 | ✓ 905ms | ✓ 1860ms | ✓ 1419ms | http |
| 34.101.184.164:3128 | ✓ 1616ms | 否 | ✓ 1235ms | ✓ 1584ms | ✓ 1059ms | http |

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
