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

最后更新：2026-04-23 16:07:36 UTC（2026-04-24 00:07:36 UTC+8）

**代理总数：50**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 790ms | ✓ 998ms | ✓ 808ms | ✓ 1046ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1436ms | 否 | ✓ 1008ms | ✓ 1169ms | ✓ 997ms | http |
| 1.231.81.166:3128 | ✓ 1686ms | 否 | 否 | ✓ 937ms | ✓ 776ms | http |
| 152.42.208.139:8118 | ✓ 765ms | 否 | ✓ 1618ms | ✓ 1349ms | ✓ 898ms | http |
| 45.140.147.82:1081 | ✓ 1303ms | 否 | ✓ 1235ms | 否 | ✓ 1694ms | http |
| 47.85.51.197:1080 | ✓ 301ms | ✓ 1238ms | ✓ 283ms | ✓ 1158ms | ✓ 882ms | http |
| 38.79.118.202:33858 | ✓ 591ms | ✓ 1594ms | ✓ 998ms | ✓ 1269ms | ✓ 897ms | http |
| 89.208.106.138:10808 | ✓ 729ms | 否 | ✓ 1682ms | 否 | ✓ 1427ms | http |
| 115.231.181.40:8128 | ✓ 907ms | ✓ 1383ms | ✓ 1290ms | 否 | 否 | http |
| 217.76.245.80:999 | ✓ 969ms | 否 | ✓ 1457ms | ✓ 1684ms | ✓ 1681ms | http |
| 35.225.22.61:80 | ✓ 1517ms | ✓ 1339ms | ✓ 1246ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1108ms | ✓ 1485ms | ✓ 1190ms | 否 | 否 | http |
| 92.113.149.172:1080 | ✓ 1175ms | 否 | ✓ 1459ms | 否 | ✓ 1789ms | http |
| 47.105.98.23:3128 | ✓ 1049ms | ✓ 1119ms | ✓ 1879ms | 否 | 否 | http |
| 58.63.109.230:10817 | ✓ 1573ms | ✓ 1275ms | 否 | ✓ 986ms | 否 | http |
| 118.113.246.123:1080 | ✓ 1261ms | ✓ 1821ms | ✓ 1436ms | 否 | 否 | http |
| 34.71.229.255:3128 | ✓ 1711ms | 否 | ✓ 1144ms | ✓ 1692ms | ✓ 884ms | http |
| 121.230.8.245:1080 | ✓ 883ms | ✓ 1068ms | ✓ 964ms | ✓ 1156ms | ✓ 955ms | http |
| 121.230.9.113:1080 | ✓ 1823ms | 否 | ✓ 1305ms | 否 | ✓ 1047ms | http |
| 101.32.244.83:8080 | 否 | ✓ 1796ms | ✓ 889ms | ✓ 1398ms | ✓ 1149ms | http |
| 121.43.196.210:8222 | ✓ 1043ms | ✓ 1038ms | ✓ 815ms | ✓ 1048ms | ✓ 882ms | http |
| 121.43.196.213:8222 | ✓ 1101ms | ✓ 983ms | ✓ 867ms | ✓ 1106ms | ✓ 904ms | http |
| 114.55.226.123:10086 | ✓ 1113ms | 否 | ✓ 1050ms | ✓ 1223ms | ✓ 1028ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1573ms | ✓ 1339ms | ✓ 1300ms | ✓ 1545ms | http |
| 152.42.177.32:8888 | ✓ 885ms | 否 | ✓ 1495ms | ✓ 1446ms | ✓ 1238ms | http |
| 120.92.212.16:8890 | ✓ 1002ms | ✓ 1355ms | ✓ 1983ms | ✓ 1426ms | ✓ 1621ms | http |
| 212.58.132.5:8888 | ✓ 1507ms | 否 | 否 | ✓ 1659ms | ✓ 1284ms | http |
| 152.32.132.190:7890 | ✓ 1295ms | ✓ 888ms | 否 | ✓ 1157ms | ✓ 659ms | http |
| 84.47.150.125:1080 | ✓ 876ms | 否 | ✓ 1431ms | 否 | ✓ 1893ms | http |
| 130.61.174.200:1080 | ✓ 1884ms | ✓ 1957ms | 否 | ✓ 1748ms | 否 | http |
| 202.129.206.239:3128 | ✓ 1064ms | 否 | ✓ 1233ms | ✓ 1580ms | ✓ 1595ms | http |
| 120.92.108.86:7890 | ✓ 1755ms | 否 | 否 | ✓ 1709ms | ✓ 1571ms | http |
| 103.56.115.156:7890 | ✓ 654ms | ✓ 1864ms | ✓ 945ms | ✓ 964ms | ✓ 651ms | http |
| 120.92.212.16:7890 | ✓ 1231ms | 否 | ✓ 1712ms | 否 | ✓ 1984ms | http |
| 168.110.52.228:3128 | ✓ 1178ms | 否 | ✓ 1768ms | ✓ 1007ms | ✓ 1163ms | http |
| 8.219.195.129:1080 | ✓ 847ms | 否 | ✓ 721ms | ✓ 1001ms | 否 | http |
| 92.113.149.172:8080 | ✓ 963ms | 否 | ✓ 1065ms | 否 | ✓ 1646ms | http |
| 64.188.77.26:3128 | ✓ 856ms | 否 | ✓ 1600ms | 否 | ✓ 1544ms | http |
| 47.84.73.61:1080 | ✓ 1457ms | ✓ 1687ms | ✓ 727ms | ✓ 1022ms | ✓ 807ms | http |
| 45.186.6.104:3128 | ✓ 1308ms | ✓ 1960ms | ✓ 1741ms | 否 | 否 | http |
| 8.209.238.110:47701 | ✓ 455ms | 否 | ✓ 495ms | ✓ 792ms | 否 | http |
| 64.188.77.221:3128 | ✓ 826ms | 否 | ✓ 864ms | ✓ 1891ms | ✓ 1517ms | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1829ms | ✓ 1849ms | ✓ 1099ms | http |
| 200.125.171.254:999 | ✓ 1349ms | ✓ 1742ms | 否 | ✓ 1687ms | ✓ 1384ms | http |
| 61.83.43.219:3044 | ✓ 1353ms | ✓ 1101ms | ✓ 1059ms | ✓ 938ms | ✓ 754ms | http |
| 61.52.131.172:8443 | ✓ 976ms | ✓ 1152ms | ✓ 908ms | ✓ 1144ms | ✓ 930ms | http |
| 223.84.151.86:30005 | ✓ 1772ms | ✓ 1736ms | ✓ 1548ms | ✓ 1779ms | ✓ 1693ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1910ms | ✓ 1927ms | ✓ 1726ms | http |
| 62.113.119.14:8080 | ✓ 1558ms | ✓ 1754ms | ✓ 937ms | ✓ 1663ms | ✓ 1252ms | http |
| 125.64.244.100:8889 | ✓ 1611ms | ✓ 1698ms | ✓ 1693ms | 否 | 否 | http |

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
